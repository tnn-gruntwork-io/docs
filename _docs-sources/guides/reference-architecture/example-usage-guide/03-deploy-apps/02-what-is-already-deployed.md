# What's already deployed

When Gruntwork initially deploys the Reference Architecture, we deploy the
<<<<<<< Updated upstream
[aws-sample-app](https://github.com/tnn-gruntwork-io/aws-sample-app/) into it, configured both as a frontend (i.e.,
user-facing app that returns HTML) and as a backend (i.e., an app that's only accessible internally and returns JSON).
We recommend checking out the [aws-sample-app](https://github.com/tnn-gruntwork-io/aws-sample-app/) as it is designed to
=======
[aws-sample-app](https://github.com/tnn-gruntwork-io/aws-sample-app/) into it, configured both as a frontend (i.e.,
user-facing app that returns HTML) and as a backend (i.e., an app that's only accessible internally and returns JSON).
We recommend checking out the [aws-sample-app](https://github.com/tnn-gruntwork-io/aws-sample-app/) as it is designed to
>>>>>>> Stashed changes
deploy seamlessly into the Reference Architecture and demonstrates many important patterns you may wish to follow in
your own apps, such as how to package your app using Docker or Packer, do service discovery for microservices and data
stores in a way that works in dev and prod, securely manage secrets such as database credentials and self-signed TLS
certificates, automatically apply schema migrations to a database, and so on.

However, for the purposes of this guide, we will create a much simpler app from scratch so you can see how all the
pieces fit together. Start with this simple app, and then, when you're ready, start adopting the more advanced
<<<<<<< Updated upstream
practices from [aws-sample-app](https://github.com/tnn-gruntwork-io/aws-sample-app/).
=======
practices from [aws-sample-app](https://github.com/tnn-gruntwork-io/aws-sample-app/).
>>>>>>> Stashed changes
