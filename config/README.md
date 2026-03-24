### This directory contains the runtime YAML configuration file(s).

Place your environment-specific configuration files here (e.g., `config.yaml`, `config.local.yaml`).
These files are read at startup by the config loader in `internal/config/`.

Guidelines:
- Do not commit secrets (passwords, tokens, keys) to this directory — use environment variables or a secrets manager instead.
- Keep default/example values in `config.yaml` so developers can run the service with minimal setup.
- All values are read directly from the file — no environment variable substitution is performed.
