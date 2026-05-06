# 3-node cluster on DigitalOcean

Provisions three Droplets and bootstraps them as an equal-footing Cosmos cluster.

## What it does

1. Creates 3 `digitalocean_droplet.node` (Ubuntu) and a firewall opening 22, 80, 443 tcp + 4242 udp.
2. `cosmos_remote_install` runs `get-pro.sh` on each Droplet.
3. Bootstraps node 0 via `cosmos_install` (the "founder") — Let's Encrypt + Create-mode MongoDB.
4. Creates the constellation on node 0 (`cosmos_constellation`).
5. Mints two device join blobs with `cosmos_constellation_device`, **`cosmos_node = 2`** (full server peers, not agents).
6. Bootstraps node 1 and node 2 via `cosmos_install` with `constellation_config`.
7. Creates one `cosmos_deployment` with `replicas = 3`, placed cluster-wide via `least-busy`.

After step 4, all three nodes operate as equal Cosmos servers. The "founder" label only describes which node generated the CA.

## Prerequisites

- A DigitalOcean API token (`var.do_token`).
- An SSH key already uploaded to your DO account (its fingerprint goes in `var.ssh_key_fingerprint`) and the matching private key on disk.
- A DNS A record for `var.lighthouse_hostname` pointing at node 0's `ipv4_address` (required for Let's Encrypt). Either pre-create the record, or run a two-stage apply: `terraform apply -target=digitalocean_droplet.node`, set the DNS, then full `terraform apply`.
- A Cosmos Pro licence key.

## Run it

```bash
cd examples/digitalocean-3node

terraform init
terraform apply \
  -var 'do_token=…' \
  -var 'ssh_key_fingerprint=aa:bb:cc:…' \
  -var 'ssh_private_key_path=~/.ssh/id_ed25519' \
  -var 'lighthouse_hostname=cosmos.example.com' \
  -var 'admin_password=…' \
  -var 'cosmos_licence=…'
```

Outputs: `node_ips` (list of 3) and `admin_token` (sensitive).

## Notes

- Peer nodes (1 and 2) use `mongodb_mode = "DisableUserManagement"` — the founder owns user state, replicated cluster-wide via constellation.
- Peer hostnames default to their public IPv4 with `SELFSIGNED` certs. If you want Let's Encrypt on the peers too, give each its own DNS record and switch `https_certificate_mode`.
- The `cosmos_deployment.web` lives in NATS KV and is cluster-replicated, so the `cosmos.cluster` provider could equally point at any of the three nodes.
