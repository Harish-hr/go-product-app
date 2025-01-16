Backend service developed using GoLang:
Repo location: https://github.com/Harish-hr/go-product-app

Frontend service devleoped using angular 
Repo location: https://github.com/Harish-hr/product-crud

# Instruction to run both frontend and backend app
Prerequisites
  * Docker Desktop installed and running on your local machine. Verify installation with docker --version.
Steps
  * **Clone/Download:** Clone or download the project repository to your local machine.
  * **Navigate:** Open a terminal and navigate to the project directory containing the docker-compose.yml file.
  * **Start Application:** Run docker compose up -d. This command will build and start both the backend and frontend servers in detached mode.
  * **Verify Backend:** Open a browser and navigate to http://localhost:8080. The message "Welcome to the Product Management System!!!" confirms successful backend deployment.
  * **Access Frontend:** Open a browser and navigate to http://localhost:8081/products to access the product management interface.
