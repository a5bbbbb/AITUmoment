# main.tf

# This special "null" provider is often used when you need to trigger
# local actions or manage resources that don't directly map to a cloud provider.
# It doesn't provision any infrastructure itself.
provider "null" {}

# --- Variables ---
variable "app_source_path" {
  description = "Application directory."
  type        = string
  default     = "C:\\Users\\PC\\Desktop\\uni_courses\\Advanced_Programming_1\\AITUmoment"
}


# --- Null Resource to Trigger Local Commands ---
resource "null_resource" "docker_compose_deploy" {
  triggers = {
    app_path_hash      = filebase64sha256("${var.app_source_path}/docker-compose.yml")
    destination_path   = var.app_source_path
  }

  # --- Provisioners ---
  # These 'local-exec' provisioners will run commands on your local machine.

  # 1. Stop and remove existing Docker Compose services (optional but recommended for redeployment)
  provisioner "local-exec" {
    command           = "docker compose -f ${var.app_source_path}/docker-compose.yml down || true"
  }

  # 2. Run Docker Compose
  provisioner "local-exec" {
    command           = "docker compose -f ${var.app_source_path}/docker-compose.yml up -d"
  }

  # 3. Clean up (optional - run on destroy)
  # This provisioner will execute when you run `terraform destroy`
  provisioner "local-exec" {
    when              = destroy
    command           = "docker compose -f ${self.triggers.destination_path}/docker-compose.yml down"
  }
}

# --- Outputs (optional) ---
output "deployment_path" {
  value       = var.app_source_path
}