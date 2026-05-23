# H1R4 VPN

<p align="center">
  <img src="resources\img\thunder.png" width="700">
</p>

<p align="center">
  Lightweight • Fast • Go-Powered OpenVPN Controller
</p>

---

![Go](https://img.shields.io/badge/Go-1.24+-00ADD8?style=for-the-badge&logo=go)
![Platform](https://img.shields.io/badge/Platform-Linux-black?style=for-the-badge&logo=linux)
![OpenVPN](https://img.shields.io/badge/OpenVPN-Compatible-orange?style=for-the-badge)
![Status](https://img.shields.io/badge/Status-Active-success?style=for-the-badge)

---

## ✨ Features

- ⚡ Ultra-lightweight Go controller
- 🔒 Secure OpenVPN process management
- 🧠 Smart wrapper compatibility
- 📜 Logging support
- 🌐 Public IP detection
- 🧩 Bash / Perl / Ruby wrappers
- 🐧 Linux-first architecture
- 🚀 Minimal RAM usage

---

## 📦 Installation

### 1. Clone Repository

```bash
git clone https://github.com/4NG3RS0N/h1r4vpn.git
cd h1r4vpn
```

### 2. Install OpenVPN

#### Debian / Ubuntu

```bash
sudo apt update
sudo apt install openvpn -y
```

#### Arch Linux

```bash
sudo pacman -S openvpn
```

#### Fedora

```bash
sudo dnf install openvpn
```

---

## 🔨 Build

```bash
go build -o h1r4vpn ./cmd/h1r4vpn
```

---

## 🚀 Usage

### Check Status

```bash
./h1r4vpn status
```

### Connect

```bash
sudo ./h1r4vpn connect --config profile.ovpn
```

### Disconnect

```bash
sudo ./h1r4vpn disconnect
```

---

## 📜 Wrappers

### Bash

```bash
./wrappers/h1r4vpn.sh status
```

### Perl

```bash
./wrappers/h1r4vpn.pl status
```

### Ruby

```bash
./wrappers/h1r4vpn.rb status
```

---

## 📂 Logs

Logs are stored inside:

```text
logs/h1r4vpn.log
```

---

## 🔒 Security Notes

- No shell injection
- PID-based process management
- Minimal process spawning
- OpenVPN handled directly
- Safer wrapper execution

---

## 🛣 Roadmap

- [ ] OpenVPN Management Interface
- [ ] WireGuard support
- [ ] TUI dashboard
- [ ] Live bandwidth monitor
- [ ] Multi-profile management
- [ ] DNS leak protection
- [ ] Kill switch

---

## 📜 License

No License

---

## ❤️ Author

Built with Iced-Coffee, packets, love and questionable sleep schedules by **4NG3RS0N**