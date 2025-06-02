# main.tf
provider "null" {}

variable "app_source_path" {
  description = "Application directory."
  type        = string
  default     = "C:\\Users\\PC\\Desktop\\uni_courses\\Advanced_Programming_1\\AITUmoment"
}


resource "null_resource" "docker_compose_deploy" {
  triggers = {
    app_path_hash      = filebase64sha256("${var.app_source_path}/docker-compose-hub.yml")
    destination_path   = var.app_source_path
  }

  provisioner "local-exec" {
    command           = "docker compose -f ${var.app_source_path}/docker-compose-hub.yml down || true"
  }

  provisioner "local-exec" {
    command           = "docker compose -f ${var.app_source_path}/docker-compose-hub.yml up -d"
  }

  provisioner "local-exec" {
    when              = destroy
    command           = "docker compose -f ${self.triggers.destination_path}/docker-compose-hub.yml down"
  }
}

# --- Outputs (optional) ---
output "deployment_path" {
  value       = var.app_source_path
}