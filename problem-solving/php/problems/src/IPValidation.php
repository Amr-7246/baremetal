<?php

function isValidIP(string $ip) {
    //& Insure that the IP containes 3 " . " -> explode the full id by the " . " to insure that it between 0, 255 and do not start with 0
    if (substr_count($ip, '.') != 3 ) { return false; }

    $ip_chuncks_array = explode('.', $ip);
    if (count($ip_chuncks_array ) != 4){ return false;}

    foreach ($ip_chuncks_array as $chunck) {
        if (!is_numeric($chunck) || (int)$chunck > 255 || (int)$chunck < 0 || ($chunck[0] == "0" && strlen($chunck) > 1 ) || $chunck[0] == " " || $chunck[-1] == " "){
            return false; 
        }
    }
    return true;
}
var_dump(isValidIP("1.5.2.3 "));
