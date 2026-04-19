<?php

//& The auto loader 
    //! I know I should use composer to handle that, but the goal is to leaning deeply not building easer....
    //TODO: optimaize that code comparing with composer`s findFileWithExtension function 
    //TODO: Handle the Exception so the app don't just die with a white screen if a class is missing
    spl_autoload_register(function ($class){
        $root = dirname(__DIR__) . DIRECTORY_SEPARATOR; //~ "dirname(__DIR__)" represent the folder that contain the folder of the index.php file, meaning it do "../" for the current dir to return to the project root 
        $filePath = str_replace("\\", DIRECTORY_SEPARATOR, $class) . '.php';
        $fullPath = $root . lcfirst($filePath);

        if(file_exists($fullPath)){ //~ Yup, checking the disc for each file is slow, we can follow composer trick to solve that by scaning the full project and create a classMap with its classes - "composer dump-autoloaded -o" comman - but it is very hard coding ...
            include_once $fullPath;
        }
    });