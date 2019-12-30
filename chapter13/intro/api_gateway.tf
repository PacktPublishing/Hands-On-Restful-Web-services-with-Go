resource "aws_api_gateway_rest_api" "test" {
  name        = "EC2Example"
  description = "Terraform EC2 REST API Example"

  endpoint_configuration {
    types = ["REGIONAL"]
  }
}

// Method request configuration
resource "aws_api_gateway_method" "test" {
   rest_api_id   = aws_api_gateway_rest_api.test.id
   resource_id   = aws_api_gateway_rest_api.test.root_resource_id
   http_method   = "GET"

   authorization = "NONE"
 }

// Method response configuration
resource "aws_api_gateway_method_response" "test" {
  rest_api_id = aws_api_gateway_rest_api.test.id
  resource_id = aws_api_gateway_rest_api.test.root_resource_id
  http_method = aws_api_gateway_method.test.http_method

  status_code = "200"
}

// Integration request configuration
resource "aws_api_gateway_integration" "test" {
   rest_api_id = aws_api_gateway_rest_api.test.id
   resource_id = aws_api_gateway_method.test.resource_id
   http_method = aws_api_gateway_method.test.http_method

   integration_http_method = "GET"
   type                    = "HTTP"
   uri                     = "http://${aws_instance.api_server.public_dns}/api/books"
 }

// Integration response configuration
resource "aws_api_gateway_integration_response" "MyDemoIntegrationResponse" {
  rest_api_id = aws_api_gateway_rest_api.test.id
  resource_id = aws_api_gateway_rest_api.test.root_resource_id
  http_method = aws_api_gateway_method.test.http_method

  status_code = aws_api_gateway_method_response.test.status_code
}

// Deploy API on Gateway with test environment
resource "aws_api_gateway_deployment" "test" {
   depends_on = [
     aws_api_gateway_integration.test
   ]

   rest_api_id = aws_api_gateway_rest_api.test.id
   stage_name  = "test"
 }
