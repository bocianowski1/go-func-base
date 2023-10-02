resource "azurerm_resource_group" "rg" {
  name     = "${var.prefix}-resources"
  location = local.location

  tags = {
    environment = local.environment
  }
}

module "function" {
  source = "./modules/function"

  rg_name     = azurerm_resource_group.rg.name
  location    = local.location
  prefix      = var.prefix
}
