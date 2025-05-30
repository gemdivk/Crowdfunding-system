resource "google_compute_instance" "go_app_instance" {
  name         = var.instance_name
  machine_type = "e2-medium"
  zone         = var.zone

  boot_disk {
    initialize_params {
      image = "debian-cloud/debian-11"
    }
  }

  network_interface {
    network = "default"
    access_config {}
  }

  metadata_startup_script = <<-EOT
    #!/bin/bash

    sudo apt-get update
    sudo apt-get install -y golang git postgresql postgresql-contrib

    sudo -u postgres psql -c "ALTER USER postgres PASSWORD '${var.db_password}';"
    sudo -u postgres createdb appdb


    git clone https://github.com/gemdivk/Crowdfunding-system.git /opt/app

    cd /opt/app

    go build -o app main.go

    nohup ./app &
  EOT

  tags = ["http-server"]
}
