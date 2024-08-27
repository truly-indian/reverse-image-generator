# Reverse Image Generator
- It is a backend service that powers a website which requires a reverse image search functionality.

# Tech Stack Used
  Golang
  Gin
  Google Wire

# API Details

- Reverse Image Genrator API:

  end_point: /api/v1/reverse_image_generator
  request_body: {
    "imageUrl": "some image url pointing to image",
    "pageToken": 1
  }
  response: {
    "Products": [
      {
        "productName": "name of the product",
        "price": "price of the product",
        "rating": "1"
      }
    ]
  }

# Steps To Run the Project
- Clone the repository.
- Make sure that you have go, wire installed on your system if not please install them first.
- Run this command to build
  make compile
- Run this command to run the project
  ./out/reverse-image-generator start

# Screen Shots
<img width="1391" alt="Screenshot 2024-08-27 at 8 32 45 AM" src="https://github.com/user-attachments/assets/83128e49-633c-4a83-ab54-e283a8d872bf">
