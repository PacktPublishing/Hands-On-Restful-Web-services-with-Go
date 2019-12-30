resource "aws_api_gateway_rest_api" "test" {
  name        = "EC2Example"
  description = "Terraform EC2 REST API Example"
}

// Resource is an endpoint
resource "aws_api_gateway_resource" "proxy" {
   rest_api_id = aws_api_gateway_rest_api.test.id
   parent_id   = aws_api_gateway_rest_api.test.root_resource_id
   path_part   = "{proxy+}"
}

// Method is a client HTTP method
resource "aws_api_gateway_method" "proxy" {
   rest_api_id   = aws_api_gateway_rest_api.test.id
   resource_id   = aws_api_gateway_resource.proxy.id
   http_method   = "GET"
   authorization = "NONE"
 }

// Target endpoint configuration
resource "aws_api_gateway_integration" "http_server" {
   rest_api_id = aws_api_gateway_rest_api.test.id
   resource_id = aws_api_gateway_method.proxy.resource_id
   http_method = aws_api_gateway_method.proxy.http_method

   integration_http_method = "GET"
   type                    = "AWS_PROXY"
   uri                     = aws_instance.api_server.arn
 }

// Root resource method
resource "aws_api_gateway_method" "proxy_root" {
   rest_api_id   = aws_api_gateway_rest_api.test.id
   resource_id   = aws_api_gateway_rest_api.test.root_resource_id
   http_method   = "GET"
   authorization = "NONE"
 }

// Root resource integration
 resource "aws_api_gateway_integration" "http_server_root" {
   rest_api_id = aws_api_gateway_rest_api.test.id
   resource_id = aws_api_gateway_method.proxy_root.resource_id
   http_method = aws_api_gateway_method.proxy_root.http_method

   integration_http_method = "GET"
   type                    = "AWS_PROXY"
   uri                     = aws_instance.api_server.arn
 }

resource "aws_api_gateway_deployment" "example" {
   depends_on = [
     aws_api_gateway_integration.http_server,
     aws_api_gateway_integration.http_server_root,
   ]

   rest_api_id = aws_api_gateway_rest_api.test.id
   stage_name  = "test"
 }
