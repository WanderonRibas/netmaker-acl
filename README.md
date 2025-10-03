&lt;p align="center"&gt;
  &lt;a href="https://netmaker.io "&gt;
  &lt;img src="https://raw.githubusercontent.com/gravitl/netmaker-docs/master/images/netmaker-github/netmaker-teal.png " width="50%"&gt;&lt;break/&gt;
  &lt;/a&gt;
&lt;/p&gt;

&lt;p align="center"&gt;
&lt;a href="https://runacap.com/ross-index/annual-2022/ " target="_blank" rel="noopener"&gt;
    &lt;img src="https://runacap.com/wp-content/uploads/2023/02/Annual_ROSS_badge_white_2022.svg " alt="ROSS Index - Fastest Growing Open-Source Startups | Runa Capital" width="17%" /&gt;
&lt;/a&gt;  
&lt;a href="https://www.ycombinator.com/companies/netmaker/ " target="_blank" rel="noopener"&gt;
    &lt;img src="https://raw.githubusercontent.com/gravitl/netmaker-docs/master/images/netmaker-github/y-combinator.png " alt="Y-Combinator" width="16%" /&gt;
&lt;/a&gt;  
&lt;/p&gt;

&lt;p align="center"&gt;
  &lt;a href="https://github.com/gravitl/netmaker/releases "&gt;
    &lt;img src="https://img.shields.io/badge/Version-1.1.0-informational?style=flat-square " /&gt;
  &lt;/a&gt;
  &lt;a href="https://hub.docker.com/r/gravitl/netmaker/tags "&gt;
    &lt;img src="https://img.shields.io/docker/pulls/gravitl/netmaker?label=downloads " /&gt;
  &lt;/a&gt;  
  &lt;a href="https://goreportcard.com/report/github.com/gravitl/netmaker "&gt;
    &lt;img src="https://goreportcard.com/badge/github.com/gravitl/netmaker " /&gt;
  &lt;/a&gt;
  &lt;a href="https://twitter.com/intent/follow?screen_name=netmaker_io "&gt;
    &lt;img src="https://img.shields.io/twitter/follow/netmaker_io?label=follow&style=social " /&gt;
  &lt;/a&gt;
  &lt;a href="https://www.youtube.com/channel/UCach3lJY_xBV7rGrbUSvkZQ "&gt;
    &lt;img src="https://img.shields.io/youtube/channel/views/UCach3lJY_xBV7rGrbUSvkZQ?style=social " /&gt;
  &lt;/a&gt;
  &lt;a href="https://reddit.com/r/netmaker "&gt;
    &lt;img src="https://img.shields.io/reddit/subreddit-subscribers/netmaker?label=%2Fr%2Fnetmaker&style=social " /&gt;
  &lt;/a&gt;  
  &lt;a href="https://discord.gg/zRb9Vfhk8A "&gt;
    &lt;img src="https://img.shields.io/discord/825071750290210916?color=%09%237289da&label=chat " /&gt;
  &lt;/a&gt; 
&lt;/p&gt;

# WireGuard&lt;sup&gt;®&lt;/sup&gt; automation from homelab to enterprise

| Create                                    | Manage                                  | Automate                                |
|-------------------------------------------|-----------------------------------------|-----------------------------------------|
| :heavy_check_mark: WireGuard Networks     | :heavy_check_mark: Admin UI             | :heavy_check_mark: Linux                |
| :heavy_check_mark: Remote Access Gateways | :heavy_check_mark: OAuth                | :heavy_check_mark: Docker              |
| :heavy_check_mark: Mesh VPNs              | :heavy_check_mark: Private DNS          | :heavy_check_mark: Mac                  |
| :heavy_check_mark: Site-to-Site           | :heavy_check_mark: Access Control Lists | :heavy_check_mark: Windows              |

# Try Netmaker SaaS  

If you're looking for a managed service, you can get started with just a few clicks, visit [netmaker.io](https://account.netmaker.io ) to create your netmaker server.  

# Self-Hosted Open Source Quick Start  

These are the instructions for deploying a Netmaker server on your cloud VM as quickly as possible. For more detailed instructions, visit the [Install Docs](https://docs.netmaker.io/docs/server-installation/quick-install#quick-install-script ).  

1. Get a cloud VM with Ubuntu 24.04 and a static public IP.
2. Allow inbound traffic on port 443,51821 TCP and UDP to the VM firewall in cloud security settings, and for simplicity, allow outbound on All TCP and All UDP.
3. (recommended) Prepare DNS - Set a wildcard subdomain in your DNS settings for Netmaker, e.g. *.netmaker.example.com, which points to your VM's public IP.
4. Run the script to setup open source version of Netmaker: 

`sudo wget -qO /root/nm-quick.sh https://raw.githubusercontent.com/gravitl/netmaker/master/scripts/nm-quick.sh  && sudo chmod +x /root/nm-quick.sh && sudo /root/nm-quick.sh`

**&lt;pre&gt;To Install Self-Hosted PRO Version - https://docs.netmaker.io/docs/server-installation/netmaker-professional-setup &lt;/pre&gt;** 

---

## 🔐 ACL por Rede (Network-scoped Permissions) – Fork WanderonRibas

Este fork adiciona **controle de acesso fino por rede**:  
usuários comuns só visualizam e gerenciam as redes explicitamente autorizadas.

### Como habilitar
1. Compile normalmente:
   ```bash
   go mod tidy && go build -o netmaker .
