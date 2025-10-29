# ResumeAdvisor (gRPC)

A Go gRPC service that provides **resume feedback**. Users can submit a short resume snippet and get career-oriented suggestions.

## Features
- Unary RPC: Get resume feedback
- Streaming RPC: Optional detailed line-by-line feedback
- Modular architecture: controller + service