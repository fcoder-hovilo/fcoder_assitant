# Discord Bot Project

This is a simple Discord bot project written in Go that demonstrates basic bot functionality, command handling, and periodic message scheduling.

## Table of Contents

- [Features](#features)
- [Project Structure](#project-structure)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Configuration](#configuration)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Features

- Command handling with a specified prefix.
- Basic command example: `!hello`.
- Periodic message scheduling (every 1 minute).

## Project Structure

The project is organized into the following packages:

- **main**: Contains the main application entry point.
- **config**: Handles configuration loading from a JSON file.
- **handlers**: Manages Discord event handlers, such as message creation.
- **scheduler**: Implements the periodic message scheduling logic.

## Prerequisites

- [Go](https://golang.org/doc/install) installed on your machine.
- A Discord bot token. You can obtain one by [creating a new bot on Discord](https://discord.com/developers/applications).

## Getting Started

1. Clone this repository:

   ``` bash
   git clone https://github.com/khengyun/fcoder_assitant.git
    ```
2. Navigate to the project directory:
    ``` bash
    cd fcoder_assitant
    ```
3. Install dependencies:
    ``` bash
    go mod download
    ```

## Configuration
Create a `config.json` file in the `config` directory with the following structure:
```json
{
  "Prefix": "!",
  "Token": "YOUR_DISCORD_BOT_TOKEN"
}
```
Replace `"YOUR_DISCORD_BOT_TOKEN"` with your actual bot token.

## Usage
1. Run the bot:
```bash
go run main.go
```
2. Invite the bot to your Discord server using the bot authorization URL generated when you created the bot on the Discord Developer Portal.
3. Interact with the bot using the specified prefix and commands.

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, feel free to open an issue or create a pull request.

## License
This project is licensed under the Fcoder License.