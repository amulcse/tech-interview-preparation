<?php
$string = 'abba';
$pushLimit = ceil(strlen($string)/2);
$stack = [];
$odd = strlen($string)%2 ? true : false ;
for($i=0;$i<strlen($string);$i++){
	
	if($i < $pushLimit){
		if(!($odd && $i+1 == $pushLimit)){
			array_push($stack,$string[$i]);
		}
	}
	else{
		if(array_pop($stack) !== $string[$i]){
			echo "Not palindrom";exit;
		}
	}	
}

if(sizeof($stack) > 0){
	echo "Not Palindrome";
}
else{
	echo "Palindrome";
}