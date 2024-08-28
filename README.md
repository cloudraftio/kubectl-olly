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

You can download the latest binary from the [release page](https://github.com/cloudraftio/kubectl-olly/releases).

### Install via Krew

👉 Ensure you have [kubectl](https://kubernetes.io/docs/tasks/tools/#kubectl) installed and configured.
👉 Install the plugin using:
   ```bash
   kubectl krew install olly
   ```

### Install via [HomeBrew](https://brew.sh/) on macOS/Linux

```shell
brew install
```

### Install via go

```shell
go install github.com/cloudraftio/kubectl-olly/@latest
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
