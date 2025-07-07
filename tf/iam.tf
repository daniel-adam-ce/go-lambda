resource "aws_iam_role" "iam_for_go_lambda" {
  name = "go-lambda-role"
  assume_role_policy = data.aws_iam_policy_document.go_lambda_role.json
}

resource "aws_iam_role_policy_attachment" "trigger_lambda_policy" {
  role       = aws_iam_role.iam_for_go_lambda.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole"
}
