# Priofirst-CLI

A Command Line Interface (CLI) tool for managing events. This CLI allows users to add, view, and manage events in a SQLite database.

## Features

- **Add Events**: Prompt the user to enter event details such as title, priority, alert type, due date, due time, and notes.
- **Store Events**: Automatically store events in a SQLite database.
- **Manage Events**: Manage event priorities, alerts, and details.

## Table of Contents

- [Installation](#installation)
- [Usage](#usage)
- [Database Schema](#database-schema)
- [Commands](#commands)
- [Contributing](#contributing)
- [License](#license)

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.16+)
- SQLite3

### Clone the Repository

```bash
git clone https://github.com/yourusername/events-cli.git
cd events-cli

go build -o Priofirst

prfs add-events

