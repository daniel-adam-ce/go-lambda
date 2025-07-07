data "aws_iam_policy_document" "go_lambda_role" {
    statement {
        effect = "Allow"
        principals {
            type = "Service"
            identifiers = ["lambda.amazonaws.com"]
        }
        
        actions = ["sts:AssumeRole"]
    }
}