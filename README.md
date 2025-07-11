<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="./media/3x-ui-dark.png">
    <img alt="3x-ui" src="./media/3x-ui-light.png">
  </picture>
</p>

**3X-UI** — advanced, open-source web-based control panel designed for managing Xray-core server. It offers a user-friendly interface for configuring and monitoring various VPN and proxy protocols.

> [!IMPORTANT]
> This project is only for personal using, please do not use it for illegal purposes, please do not use it in a production environment.

As an enhanced fork of the original X-UI project, 3X-UI provides improved stability, broader protocol support, and additional features.

## PostgreSQL Support

This version of 3X-UI includes PostgreSQL database support, allowing you to connect to an external PostgreSQL database server. This enables you to:

- Store your 3X-UI data on a separate database server
- Use a managed PostgreSQL service (like AWS RDS, Google Cloud SQL, etc.)
- Implement database clustering and high availability
- Separate your application and data layers

For detailed setup instructions, please refer to the [PostgreSQL Setup Guide](POSTGRESQL_SETUP.md).

## Quick Start

```bash
bash <(curl -Ls https://raw.githubusercontent.com/mordak-95/3x-ui/master/install.sh)
```
