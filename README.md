To use this project, ensure you have Golang installed, preferably version 1.22.2 or newer.
To check it run 
```
go version
```
Then navigate to the 'main' directory, to execute a demo:
```
cd main
```
Then run it with following command:
```
go run .
```
This will execute the demo. The example JSON file, 'aws_iam_role.json', will be read automatically. Feel free to modify it according to your needs.

The main purpose of this project is the 'jsonanalyze' module, which provides all the necessary tools to verify input JSON data and check if it fulfills the requirements of an AWS::IAM::Role Policy. It's a ready-to-use module, so feel free to incorporate it into your application!

To test the 'jsonanalyze' module execute this commands:
```
cd jsonanalyze
go test
```
To see specific tests results run:
```
go test -v
```

For more information about the format of the AWS::IAM::Role Policy, you can refer to the following documentation:
* [AWS::IAM::Role Policy](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-role-policy.html)
* [Overview of JSON policies](https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html#access_policies-json)
* [IAM JSON policy elements reference](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_policies_elements.html)
* [IAM identifiers](https://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-arns)

Author: [Adam Mytnik](mailto:adammytnik@student.agh.edu.pl)
