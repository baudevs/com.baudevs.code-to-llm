# ctllm - Code to LLM

`ctllm` (Code to LLM) is a CLI tool designed to simplify the process of preparing codebases for Large Language Models (LLMs). It recursively scans your project directory, collects relevant code files, and organizes them into manageable chunks, respecting token limits. The tool applies custom ignore patterns, integrates with `.gitignore`, and provides an interactive initialization process to tailor settings to your project.

## Features

- **Recursive File Collection**: Automatically traverses your project directory to gather all code files.
- **Custom Ignore Patterns**: Supports custom ignore patterns via an `ignore_patterns.yaml` file.
- **.gitignore Integration**: Reads and applies patterns from your `.gitignore` file.
- **Token Limit Handling**: Splits files into chunks based on a specified token limit for optimal LLM processing.
- **Interactive Initialization**: Provides an interactive setup to configure your project, including output directory and ignore patterns.
- **Project Type Detection**: Automatically detects your project's type (e.g., Node.js, Python, Go) and suggests default ignore patterns.
- **Output Directory Management**: Optionally adds the output directory to your `.gitignore` file to prevent committing generated files.

## Installation

To install `ctllm`, follow these steps:

### Option 1: Install cloning the repository

1. **Clone the Repository**:

   ```bash
   git clone https://github.com/baudevs/code-to-llm.git
   ```

2. **Navigate to the Project Directory**:

```bash
cd ctllm
```

3. **Install the Tool**:

Ensure you have Go installed on your system.
To check if you have go installed, run:

```bash
go version
```

if you don't see the version, or an error it means that you don't have go installed.
you can install it from [https://golang.org/dl/](https://golang.org/dl/).

Once you have go installed, you can install `ctllm` by running (cd to the cloned repository directory, probably com.baudevs.code-to-llm):

```bash
chmod +x install.sh
./install.sh
```

4. **Reload your shell**:

```bash
source ~/.bashrc
``` 

or if you are using zsh:

```zsh
source ~/.zshrc
``` 

5. **Test the installation**:

```bash
ctllm --help
```

### Option 2: Install the pre-compiled binary

1. **Download the Binary**:

   Visit the [Releases](https://github.com/baudevs/code-to-llm/releases) page to download the latest pre-compiled binary for your operating system.

2. **Make the Binary Executable**:

```bash
chmod +x ./install.sh
```


3. **Run the installation script**:

```bash
./install.sh
```

- You may be prompted for your password to install the executable to /usr/local/bin.

## Option 3: Manual installation

1. **Download the appropriate binary for your operating system**:

   Visit the [Releases](https://github.com/baudevs/code-to-llm/releases) page to download the latest pre-compiled binary for your operating system.

2. **Make the Binary Executable**:  

```bash
chmod +x ctllm
```

3. **Move the Binary to a Directory in Your PATH**:

```bash
sudo mv ctllm /usr/local/bin/ # or to another directory in your PATH
```

- You may be prompted for your password to install the executable to /usr/local/bin.

4. **Test the installation**:

```bash
ctllm --help
```

## Usage

### Initialize a project

```bash
ctllm init
```

This command will:

- Prompt you for the root directory of your project.
- Ask for the output directory where processed files will be stored.
- Detect your project type and allow you to confirm or change it.
- Load ignore patterns from ignore_patterns.yaml based on your project type.
- Offer to add the output directory to your .gitignore file.

Re-initialize with Force Option

To re-initialize an already initialized project:

```bash
ctllm init --force
```

This command will re-initialize the project, overwriting the existing configuration after prompting for confirmation.

### Process a project

```bash
ctllm   
```

This command will:

- Collect all relevant files, applying the ignore patterns.
- Generate a tree structure of your project and save it as file_tree.txt in the output directory.
- Split the collected files into chunks based on the token limit specified during initialization.
- Save the chunks as code_chunk_1.txt, code_chunk_2.txt, etc., in the output directory.

Sync Configuration

If you’ve made changes to your project structure or ignore_patterns.yaml, synchronize your configuration:

```bash
ctllm sync
```

This command will:

- Update the ignore patterns in the configuration.
- Re-process the project files to reflect any changes in the ignore patterns.
- Save the updated chunks as code_chunk_1.txt, code_chunk_2.txt, etc., in the output directory.

Configuration

ctllm-config.yaml

This file stores your project configuration:

- root: The root directory of your project.
- output_dir: The directory where processed files are saved.
- token_limit: The maximum token limit per file chunk.
- project_type: The type of your project.
- ignore_patterns: A list of patterns to ignore when collecting files.

ignore_patterns.yaml

Customize ignore patterns for different project types. Example:

```yaml
common:
  - .DS_Store
  - LICENSE
  - README.md
  - .git/
  - .gitignore

Node.js + Express:
  - node_modules/
  - dist/
  - build/
  - .env
  - npm-debug.log
  - yarn-error.log

Python with Flask:
  - venv/
  - __pycache__/
  - *.pyc
  - .env
```

To add more or remove project types, ignore patterns or common patterns, you can edit the local configuration file ***ctllm-config.yaml*** created during the initialization process.
This ignore patterns file is used only during initialization to generate common ignore patterns based on the project type and should not be edited by the user.

> If you have more patterns that would like to have included by default for all projects, you can create a new Pull Request by forking the repository and changing the file ***ignore_patterns.yaml***. then you can push your changes to the main repository and create a new PR.

## Contributing

We welcome contributions to `ctllm`! If you have suggestions or improvements, please open an issue or submit a pull request (Fork the repository and create a new branch with your changes).

```bash
git checkout -b feature/my-feature
```

For example, if you want to add a new project type and its corresponding ignore patterns, you can do so by adding a new entry in the ***ignore_patterns.yaml*** file.

1. Fork the Repository: Click the “Fork” button on GitHub.

```bash
git clone https://github.com/your-username/code-to-llm.git
cd code-to-llm
git checkout -b feature/new-ignore-patterns
```

2. Make your changes.

3. Commit your changes.

```bash
git commit -m "Add new project type and ignore patterns"
git push origin feature/new-ignore-patterns
```

4. Create a Pull Request.

5. Wait for the PR to be reviewed and merged.

## Licensing

This project is licensed under a dual license model:

1. MIT License: For open-source, personal, or internal projects. See the LICENSE file for details.
2. Commercial License: Required for commercial use, such as production environments or providing services to customers. See the COMMERCIAL_LICENSE file for details.

- Includes support, warranties, and custom agreements.

License Summary

- Non-commercial use: Free under the MIT License.
- Commercial use: Requires a commercial license.

For commercial licensing, contact [licenses@baudevs.com].

Support

- Open-source users: Submit issues or contribute via pull requests.
- Commercial users: Receive direct support as part of the commercial license.

## Credits

Developed to streamline the process of preparing codebases for Large Language Models (LLMs).

- [go-gitignore](https://github.com/sabhiram/go-gitignore) for the gitignore implementation.
- [promptui](https://github.com/manifoldco/promptui) for the promptui implementation.
- [gitignore](https://pkg.go.dev/github.com/sabhiram/go-gitignore) for the gitignore implementation.

Visit our GitHub page at [baudevs](https://github.com/baudevs)

Visit our website at [baudevs](https://baudevs.com)
