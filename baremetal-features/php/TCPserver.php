<?php
//& init the socket
    $sock = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);
    socket_bind($sock,"127.0.0.1",10000);
    socket_listen($sock);

//& start the listing loop, accept the connections, captcher the incoming data and termenate connections to freeup resources
do {
    $sockMsg =  socket_accept($sock);
    $msg = "\nHi there. \n" .
        "To quit, just type 'quit'. To shut down the server type 'shutdown'.\n";
    socket_write($sockMsg, $msg);
    while (true) {
        $input = socket_read($sockMsg,);
        if(!$input = trim($input)){continue;} //! explain that line
        if($input == 'quit'){break;}
        if($input == 'shutdown'){socket_close($sockMsg); break 2;} //! what does 2 do
        $talkback = "PHP: You said '$buf'.\n";
        socket_write($msgsock, $talkback, strlen($talkback));
        echo "$buf\n";
    }
    socket_close($sockMsg);
} while (true);

socket_close($sock);

// set_time_limit(0);

// $host = "0.0.0.0";
// $port = 8080;

// $socket = socket_create(AF_INET, SOCK_STREAM, SOL_TCP);
// if ($socket === false) {
//     die("Socket creation failed: " . socket_strerror(socket_last_error()));
// }

// socket_bind($socket, $host, $port);
// socket_listen($socket);

// echo "Server listening on port $port\n";

// $clients = [];

// while (true) {
//     $readSockets = $clients;
//     $readSockets[] = $socket;

//     $write = $except = null;

//     socket_select($readSockets, $write, $except, null);

//     if (in_array($socket, $readSockets)) {
//         $newClient = socket_accept($socket);
//         $clients[] = $newClient;

//         echo "New client connected\n";
//         unset($readSockets[array_search($socket, $readSockets)]);
//     }

//     foreach ($readSockets as $client) {
//         $data = @socket_read($client, 1024, PHP_NORMAL_READ);

//         if ($data === false) {
//             $index = array_search($client, $clients);
//             socket_close($client);
//             unset($clients[$index]);
//             echo "Client disconnected\n";
//             continue;
//         }

//         $data = trim($data);

//         if ($data === "PING") {
//             $response = "PONG\n";
//         } elseif ($data === "TIME") {
//             $response = date("Y-m-d H:i:s") . "\n";
//         } else {
//             $response = "Unknown command\n";
//         }

//         socket_write($client, $response);
//     }
// }