# Upgrading

This document describes breaking changes and manual actions required on major upgrades.

## v1.1.0 to v2.0.0

The data directory has moved to the standard application data directory. To upgrade, follow the instructions here.

Stop the software by closing the console window and rename the `data/gorm.db` file to `envelope-zero.db`

Now, move the file to the directories described below depending on your operating system.

- on Windows: `%APPDATA%/envelope-zero`
- on macOS: `~/Library/Application Support/envelope-zero`
- on all other Unix based systems: `~/.local/share/envelope-zero`

You will need to create the `envelope-zero` directory before copying the file into it.
