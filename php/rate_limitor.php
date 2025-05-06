<?php

/* 
A rate limiter is a crucial component for controlling traffic, preventing abuse, and ensuring fair resource usage. Here’s how you can design one:

    1. Define Requirements
    Limit Type: Requests per second, per minute, or per user.
    
    Scope: Global, per IP, per API key, per user.
    
    Response Strategy: Reject excess requests (429 Too Many Requests) or delay processing.
    
    2. Choose a Rate Limiting Algorithm
    Several algorithms help implement rate limiting efficiently:
    
    Token Bucket (Best for Bursts)
    Each request consumes a token.
    
    Tokens refill at a fixed rate.
    
    If no tokens are available, the request is denied or delayed.
    
    Leaky Bucket (Smooth Flow)
    Requests are queued and processed at a fixed rate.
    
    Prevents sudden spikes from overwhelming the system.
    
    Fixed Window Counter (Simple)
    Counts requests in a fixed time window (e.g., 60 requests per minute).
    
    If the limit is exceeded, all extra requests are denied until the window resets.
    
    Sliding Window (More Precise)
    Uses timestamps to track requests in a rolling window.
    
    Provides better accuracy compared to the fixed window.
    
    3. Implementation Choices
    In-Memory: Fastest, but resets on server restart (e.g., Redis, local memory).
    
    Distributed: Needed for multi-node systems (e.g., Redis, API gateways).
 */

 Class RateLimiter {

    private $limit;
    private $window;
    private $key;
    private $storage = [];

    public function __construct($key, $limit, $window)
    {
        $this->key = $key;
        $this->limit = $limit;
        $this->window = $window;
    }

    /* Fixed Window Counter Algorithm */
    public function fixedWindow(){

        if(!isset($this->storage[$this->key])){
            $this->storage[$this->key]['time'] = time() ;
            $this->storage[$this->key]['count'] = 0;
        }

        $duration = time() -  $this->storage[$this->key]['time'];
        if($duration > $this->window) {
            $this->storage[$this->key]['time'] = time() ;
            $this->storage[$this->key]['count'] = 0;
        }

        if($this->storage[$this->key]['count'] < $this->limit){
            $this->storage[$this->key]['count']++;
            return true;
        }

        return false;
    }

    /* Sliding Window Algorithm */
    public function slidingWindow(){

        if(!isset($this->storage[$this->key])){
            $this->storage[$this->key] = [];
        }

        $now = time();
        $currentWindow = $now - $this->window;
        
        $this->storage[$this->key] = array_filter($this->storage[$this->key], fn($timestamp) => $timestamp > $currentWindow);

        if(count($this->storage[$this->key]) < $this->limit){
            $this->storage[$this->key][] = $now;
            return true;
        }

        return false;


    }

 }

 $limit = 5;
 $window = 60; // 60 seconds
 $key = '123.123.1231.23';
 $rateLimiter = new RateLimiter($key, $limit, $window);

for($i=0;$i<10;$i++){
    $result = $rateLimiter->slidingWindow();
    echo " ".$result;
}