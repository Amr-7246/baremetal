<?php
class Database {
    private static $instance = null; //? why it static
    private $connection;
    public function __construct() {
        $host = 'localhost';
        $dbname = 'chat_app';
        $username = 'root'; 
        $password = '';     
        $dsn = "mysql:host=$host;dbname=$dbname;charset=utf8mb4";
        $options = [
            PDO::ATTR_ERRMODE => PDO::ERRMODE_EXCEPTION,
            PDO::ATTR_DEFAULT_FETCH_MODE => PDO::FETCH_ASSOC, //! to return associative array.
            PDO::ATTR_EMULATE_PREPARES => false //! more secure against SQL injection, if set to true everything retuned is string
        ];
        try {
            $this->connection = new PDO($dsn, $username, $password, $options);
        } catch (PDOException $e) {
            die("Connection failed: " . $e->getMessage());
        }
    }
    public static function getInstance(){
        if (self::$instance == null) {
            self::$instance = new self();
        }
        return self::$instance;
    }
    public function getConnection(){ //? why do not I just make the connection variable public, why should I return it from function
        return $this->connection;
    }
}

function getDB() { //? what is that function role + why not just (new Database())->getConnection(); 
    return Database::getInstance()->getConnection(); 
}

function escape($string) { //! to prevent Cross-Site Scripting (XSS) attacks
    return htmlspecialchars($string, ENT_QUOTES | ENT_HTML5, 'UTF-8'); 
}