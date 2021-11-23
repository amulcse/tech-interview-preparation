<?php

$stringA = 'pale';
$stringB = 'ple';



function isOneAway($stringA, $stringB){
  if(strlen($stringA) - strlen($stringB) > 1 ) {
    return "No";
  }
  $editCounter = 0;
  $j = 0;
  for($i = 0; $i < max(strlen($stringB),strlen($stringA)); $i++){
    if($stringA[$i] !== $stringB[$j]){

    }
  }
}