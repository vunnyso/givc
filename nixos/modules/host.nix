{
  config,
  pkgs,
  lib,
  ...
}: let
  cfg = config.givc.host;
  givc-agent = pkgs.callPackage ../packages/givc-agent.nix {};
in
  with lib; {

    options.givc.host = {
      enable = mkEnableOption "Enable givc-host module.";

      services = mkOption {
        description = ''
        List of systemd services for the manager to administrate. Expects a space separated list.
        Should be a unit file of type 'service' or 'target'.
        '';
        type = types.listOf types.str;
        default = [""];
        example = "['my-service.service' 'poweroff.target']";
      };

      addr = mkOption {
        description = "Address (ip)";
        type = types.str;
        default = "127.0.0.1";
      };

      port = mkOption {
        description = "Port";
        type = types.str;
        default = "9000";
      };

      protocol = mkOption {
        description = "Protocol";
        type = types.str;
        default = "tcp";
      };

      admin = mkOption {
        description = ''Admin server configuration.
        '';
        type = with types; submodule {
          options = {
            enable = mkEnableOption "Admin module";
            name = mkOption {
              description = "Hostname of admin server";
              type = types.str;
            };

            addr = mkOption {
              description = "Address of admin server";
              type = types.str;
            };

            port = mkOption {
              description = "Port of admin server";
              type = types.str;
            };

            protocol = mkOption {
              description = "Protocol of admin server";
              type = types.str;
            };
          };
        };
        default = {
          enable = true;
          name = "localhost";
          addr = "127.0.0.1";
          port = "9001";
          protocol = "tcp";
        };
      };

      tls = mkOption {
        description = ''TLS options for gRPC connections.
        It is enabled by default to discourage unprotected connections, and requires paths to certificates and key being set.
        If you really want to disable it use 'tls.enable = false;' .
        '';
        type = with types; submodule {
          options = {
            enable = mkEnableOption "Enable TLS";
            caCertPath = mkOption {
              description = "Path to the CA certificate file.";
              type = str;
            };
            certPath = mkOption {
              description = "Path to the service certificate file.";
              type = str;
            };
            keyPath = mkOption {
              description = "Path to the service key file.";
              type = str;
            };
          };
        };
        default = {
          enable = true;
          caCertPath = "";
          certPath = "";
          keyPath = "";
        };
      };

    };

    config = mkIf cfg.enable {

      assertions = [
        {
          assertion = cfg.services != [];
          message = "A list of services (or targets) is required for this module to run.";
        }
        {
          assertion = !(cfg.tls.enable && (
            cfg.tls.caCertPath == "" || cfg.tls.certPath == "" || cfg.tls.keyPath == ""
          ));
          message = "The TLS option requires paths' to CA certificate, service certificate, and service key.";
        }
      ];

      systemd.services.givc-host = {
        description = "GIVC remote service manager for the host.";
        enable = true;
        after = [ "network-online.target" ];
        wantedBy = [ "multi-user.target" ];
        # @TODO add requirement for admin to run before host module
        serviceConfig = {
          Type = "exec";
          ExecStart = "${givc-agent}/bin/givc-agent";
          Restart = "always";
          RestartSec = 1;
        };
        environment = {
          "NAME" = "host";
          "ADDR" = "${cfg.addr}";
          "PORT" = "${cfg.port}";
          "PROTO" = "${cfg.protocol}";
          "TLS" = "${trivial.boolToString cfg.tls.enable}";
          "SERVICES" = "${concatStringsSep " " cfg.services}";
          "ADMIN_SERVER_NAME" = "admin";
          "ADMIN_SERVER_ADDR" = "${cfg.admin.addr}";
          "ADMIN_SERVER_PORT" = "${cfg.admin.port}";
          "ADMIN_SERVER_PROTO" = "${cfg.admin.protocol}";
        } // attrsets.optionalAttrs cfg.tls.enable {
          "CA_CERT" = "${cfg.tls.caCertPath}";
          "HOST_CERT" = "${cfg.tls.certPath}";
          "HOST_KEY" = "${cfg.tls.keyPath}";
        };
      };
      networking.firewall.allowedTCPPorts =
        let
          port = lib.strings.toInt(cfg.port);
        in
          [port];
    };

}