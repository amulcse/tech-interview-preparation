<?php

$url = "Mr John Smith   ";

echo "$url == ".URLify(trim($url));

function URLify($string){
  $newString = "";
  for($i = 0; $i < strlen($string); $i++){
    if($string[$i] === ' '){
      $newString .= '%20';
    }
    else{
      $newString .= $string[$i];
    }
  }
  return $newString;
}