# Github CLI Tool
A command-line interface (CLI) tool built with Golang to interact with GitHub repositories efficiently. This tool leverages Cobra for command handling and Resty for making API requests.


## 🚀 Features 

<ul>
  <li>List user repositories</li>
  <li>Fetch repository details</li>
  <li>Create new repositories</li>
  <li>Clone repositories</li>
  <li>Delete repositories</li>
</ul>


## 🛠 Installation 
Make sure you have Go installed on your system. Then, clone the repository and build the CLI tool:

```
# Clone the repository
git clone https://github.com/yourusername/github-cli.git
cd github-cli

# Build the binary
go build -o github-tool .

# Move to a directory in your PATH (optional)
sudo mv github-tool /usr/local/bin/ (linux/macOS)
```
After that, you need to make your compiled binary executable by running chmod:

```
chmod +x /usr/local/bin/github-tool
```

## 🔑 Configuration
This tool requires a GitHub personal access token for authentication. You can set up your token as an environment variable like this:
```
export GITHUB_TOKEN="your_personal_access_token"
```

Or you can create a .env file in the same path as your compiled code (recommended):

```
echo "GITHUB_TOKEN=your_actual_token_here" > .env
```

## 🏃 Usage
Run the CLI tool with the available commands:

```
github-tool --help
```
### 📖 Example Commands
```
github-tool list
```
```
github-tool create <repo-name> --private (optional) --desc "" (optional)
```
```
github-tool delete <repo-name>
```


## 📦 Dependencies 
<p>🐍 <a href="https://github.com/spf13/cobra" >Cobra</a> - CLI framework</p>
<p>⚡ <a href="https://github.com/go-resty/resty" style="text-decoration:none;">Resty</a> - HTTP client for API requests</p>

