<?php

function rgb($r, $g, $b){
    //& -> into a loop insuer that the num between 0, 255 then extract the qoutient/reminder 
    //& -> declare an object with the latters for the numbers between 10 and 16 
    //& -> atatch the result and return it 

    $nums = [$r, $g, $b];
    $result = '';
    $hexaLatters = [
        10 => 'A',
        11 => 'B',
        12 => 'C',
        13 => 'D',
        14 => 'E',
        15 => 'F',
    ];
    foreach($nums as $num){
        $num = $num > 255 ? 255 : $num;
        $num = $num < 0 ? 0 : $num;
        $qoutient = intdiv($num, 16);
        $reminder = ($num/16) - $qoutient;
        $reminderPro = $reminder * 16;

        if ($qoutient < 10 ){
            $result .= $qoutient;
        }elseif ($qoutient >= 10 && $qoutient < 16) {
            $result .= $hexaLatters[$qoutient];
        }else{
            print('the hxa num must be btween 0, 16');
        }

        if ($reminderPro < 10 ){
            $result .= $reminderPro;
        }elseif ($reminderPro >= 10 && $reminderPro < 16) {
            $result .= $hexaLatters[$reminderPro];
        }else{
            print('the hxa num must be btween 0, 16');
        }
        }
    return $result;
}

function rgb_V2($r, $g, $b){
    $hexaNums = '0123456789ABCDEF';
    $result = "";
    
    foreach ([$r, $g, $b] as $num) {
        $num = max(0, min($num, 255));
        $qoutient = intdiv($num, 16);
        $reminder = $num % 16;
        $result .= $hexaNums[$qoutient] . $hexaNums[$reminder];
    }
    return $result;
}

function rgb_V3($r, $g, $b){
    return dechex(max(0, min($r, 255))) . dechex(max(0, min($g, 255))) . dechex(max(0, min($b, 255)));
}

function rgb_V4($r, $g, $b){
    return sprintf("%02X%02X%02X", max(0, min($r, 255)), max(0, min($g, 255)), max(0, min($b, 255)));
}

function rgb_V5($r, $g, $b){
    return vsprintf("%02X%02X%02X",array_map(fn($n)=>max(0,min($n, 255)),[$r, $g, $b]));
}

var_dump(rgb_V5(-35, 555, 155));

