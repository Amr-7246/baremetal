<?php

session_start();
header('Content-Type : application/json');

if (!isset($_SESSION['user_id'])) {
    http_response_code(401);
    echo json_encode(['error' => 'Not authenticated']);
    exit;
}
require_once '../config/database.php';

try {
    $db = getDB();
    
    $stmt = $db->prepare("
        SELECT id, user_id, username, message, created_at 
        FROM messages 
        ORDER BY created_at DESC 
        LIMIT 50
    ");
    $stmt->execute();
    $messages = $stmt->fetchAll();
    
    $messages = array_reverse($messages);
    
    echo json_encode($messages);
    
} catch (PDOException $e) {
    http_response_code(500);
    echo json_encode(['error' => 'Database error: ' . $e->getMessage()]);
}