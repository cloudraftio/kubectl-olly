# Kubectl Plugin for Olly - The Observability Chatbot

Kubectl-olly is a kubectl plugin that integrates Olly functionality directly into your kubectl workflow. This plugin brings the power of Olly to your command line, enhancing your Kubernetes observability experience.

Welcome to **Olly**, your go-to AI chatbot for all things Observability. Whether you're diving into **Thanos**, configuring **Prometheus**, exploring **OpenTelemetry**, or setting up **Grafana** and **Grafana Mimir**, Olly is here to assist you. Trained extensively on these topics, Olly is not just a knowledge base but also an excellent programmer, ready to provide insightful, accurate answers and even help you with coding tasks related to these technologies.

## Benefits

With kubectl-olly, you can:

👉 Quickly access Olly's knowledge and capabilities directly from your terminal

👉 Get instant help with observability stack configurations and troubleshooting

👉 Generate and analyze observability-related Kubernetes manifests

This plugin bridges the gap between Kubernetes operations and observability expertise, making it easier than ever to manage and optimize your monitoring and observability setup.

## Installation

### Local Installation

👉 Download the latest binary for your operating system from the [release page](https://github.com/cloudraftio/kubectl-olly/releases).

👉 Make the binary executable:
   ```bash
   chmod +x kubectl-olly
   ```

👉 Move the binary to a directory in your PATH. For example:
   ```bash
   mv kubectl-olly /usr/local/bin/
   ```

👉 For Mac users: Before executing the file, run the following command to resolve potential security warnings:
   ```bash
   xattr -d com.apple.quarantine kubectl-olly
   ```

👉 Set the OLLY_API_KEY environment variable:
    ```bash
    export OLLY_API_KEY=fake
    ```
    You can add this line to your shell configuration file (e.g., ~/.bashrc, ~/.zshrc) to make it persistent across terminal sessions.

👉 Execute with one step:
    ```bash
    kubectl olly
    ```
    or
    ```bash
    kubectl-olly
    ```

### Install via Krew

👉 Ensure you have [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl) installed and configured.

👉 Install the plugin using [Krew](https://krew.sigs.k8s.io/) and add kubectl-olly custom index:
   ```bash
   kubectl krew index add olly https://github.com/cloudraftio/kubectl-olly.git
   ```
   ```bash
   kubectl krew install olly/olly
   ```

👉 Execute with one step:
    ```bash
    kubectl olly
    ```

👉 Set the OLLY_API_KEY environment variable:
    ```bash
    export OLLY_API_KEY=fake
    ```
    You can add this line to your shell configuration file (e.g., ~/.bashrc, ~/.zshrc) to make it persistent across terminal sessions.

👉 To Remove kubectl-olly custom index:
    ```bash
    kubectl krew index remove olly
    ```

## 🚊 Usage

To initiate an interactive session with Olly, execute the following command in your terminal:

```bash
kubectl olly
```

## 🙋‍♂️ Getting Help

We are here to help!

👉 For feature requests and bugs, file an [issue](https://github.com/cloudraftio/kubectl-olly/issues).

👉 To get notified on updates ⭐️ [star this repository](https://github.com/cloudraftio/kubectl-olly/stargazers).

## ➕ Contributing

Contributions are welcome! Please feel free to submit a Pull Request. Thanks for your interest in contributing to Cloudraft!

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
