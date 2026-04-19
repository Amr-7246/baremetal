## Description
    A lightweight, custom-built PHP framework designed to demonstrate the MVC (Model-View-Controller) pattern, PSR-4 compliant autoloading, and a centralized Routing system. The goal is to move away from "spaghetti" scripts and create a reusable engine that handles logic, data, and presentation independently.
## Task Flow
    - Environment & Autoloader (PSR-4): Set up a standard directory structure and implement a custom Autoloader class using spl_autoload_register. This ensures that when you call new Controller(), PHP knows exactly which file to load without manual require statements.
    - The Request Object: Create a class to encapsulate all superglobals ($_GET, $_POST, $_SERVER). This abstracts the global state into an object-oriented format, making it easier to test and manipulate data.
    - The Router Engine: Build a router that parses the REQUEST_URI. It should use a collection of routes to map a URL (e.g., /home) to a specific Controller class and a Method (e.g., index).
    - Base Controller & View Engine: Create a Base Controller class that provides a render() method. This method will extract variables and include PHP/HTML files from the views directory, serving as your template engine.
    - Database Wrapper (PDO): Implement a Singleton or a static Wrapper class for PDO. This should handle your connection settings and provide a clean method for executing Prepared Statements to prevent SQL injection.
    - Base Model: Create an abstract Model class that your database entities (like User or Post) will extend. This layer handles the interaction between your objects and the database tables.
    - The Front Controller (index.php): The single entry point. It initializes the Autoloader, captures the Request, passes it to the Router, and dispatches the logic to the appropriate Controller.
## Project Structure
    vanilla-framework/
    ├── app/
    │   ├── Controllers/    # Application logic (e.g., HomeController.php)
    │   ├── Models/         # Database logic and entities
    │   └── Views/          # HTML templates and PHP view files
    ├── core/               # The "Engine" (The Framework itself)
    │   ├── Router.php      # URL parsing and dispatching
    │   ├── Request.php     # GET/POST/Server data abstraction
    │   ├── Controller.php  # Base controller with render() method
    │   ├── Model.php       # Base model for database interaction
    │   └── Database.php    # PDO connection wrapper
    ├── public/             # Only folder accessible to the web
    │   ├── .htaccess       # Rewrites all URLs to index.php
    │   └── index.php       # The "Front Controller" (Entry Point)
    ├── config/             # Database credentials and app constants
    ├── storage/            # Logs, uploads, or cache
    └── README.md           # Documentation of your routing logic