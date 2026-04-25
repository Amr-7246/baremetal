<?php

session_start(); //? Exploit that context to teach me more about the sessions 
header('Content-Type : application/json'); //? what is that line role

require_once '../config/database.php';

if ($_SERVER['REQUEST_METHOD'] != 'POST') {
    http_response_code(405); //? Why I do not return anything to the client, just echo at the terminal with status code
    echo json_encode(['error' => 'Method not allowed']);
    exit;
}
//& Extract the coming data
    $input = json_decode(file_get_contents('php://input')); //? is that the point where I grasp the client data, if yes who is the php://input file
    $username = trim($input['username'] ?? '');

//& Validate data
    if (empty($username)) {
        http_response_code(400);
        echo json_encode(['error' => 'Username is required']);
        exit;
    }

    if (strlen($username) < 2) {
        http_response_code(400);
        echo json_encode(['error' => 'Username must be at least 2 characters']);
        exit;
    }

    if (strlen($username) > 50) {
        http_response_code(400);
        echo json_encode(['error' => 'Username too long (max 50 characters)']);
        exit;
    }
    if (!preg_match('/^[a-zA-Z0-9_]+$/', $username)) { //TODO: learn the regular exprestions
        http_response_code(400);
        echo json_encode(['error' => 'Username can only contain letters, numbers, and underscores']);
        exit;
    }

//& inject data at the DB
    try {
        $db = getDB();
        //~ find or create a user
            $stmt = $db->prepare('SELECT id, username FROM users WHERE username = ?');
            $stmt->execute([$username]);
            $user = $stmt->fetch();

            if(!$user){
                $stmt = $db->prepare('INSERT INTO users (username) VALUES (?)');
                $stmt->execute([$username]);
                $userId = $db->lastInsertId();
            } else {
                $userId = $user['id'];
            }

        //~ store user data in session + regenerate session id
            $_SESSION['user_id'] = $userId;
            $_SESSION['username'] = $username;
            session_regenerate_id(true); //? why that line 

        echo json_encode([ //? why just echo, where is the return
            'success' => true,
            'user' => [
                'id' => $userId,
                'username' => $username
            ]
        ]);
    } catch (PDOException $e) {
        http_response_code(500);
        echo json_encode(['error' => 'Database error: ' . $e->getMessage()]);
    }
