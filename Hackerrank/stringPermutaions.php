<?php

permutaions('amul','');

function permutaions($str,$prefix){
	
	if(strlen($str) == 0){
		echo $prefix."\n";
	}
	
	for($i = 0; $i < strlen($str) ; $i++){
		$rem = substr($str,0,$i).substr($str,$i+1);
		permutaions($rem,$prefix.$str[$i]);
	}

}