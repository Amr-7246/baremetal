<?php
namespace Core;

//TODO: build the Router Engine and create the first Route, Make the Router look at that info and decide: "Okay, the path is /contact... I should call the ContactController." 
class Request
{
    public function getPath(){
        $path = $_SERVER['REQUEST_URI'] ?? '/';
        $paramPosition = strpos($path, '?');
        if($paramPosition === false){ return $path; }
        return substr($path, 0, $paramPosition);
    }

    public function getMethod(){
        return strtolower($_SERVER['REQUEST_METHOD']);
    }

    public function getBody(){
        $body = [];
        if($this->getMethod() == 'get'){
            foreach ($_GET as $key => $value) {
                //~ "FILTER_SANITIZE_SPECIAL_CHARS" flag option is a security measure. It escapes HTML characters (like < and >) to prevent any script injection like <script>alert(1)</script>
                $body[$key] = filter_input(INPUT_GET, $key, FILTER_SANITIZE_SPECIAL_CHARS);
            }
        }
            if($this->getMethod() == 'post'){
                foreach ($_POST as $key => $value) {
                    $body[$key] = filter_input(INPUT_POST, $key, FILTER_SANITIZE_SPECIAL_CHARS);
                }
            }

        return $body;
    }
} 