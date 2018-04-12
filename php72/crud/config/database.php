<?php

return [
    'default' => 'mongodb',
    'connections' => [
        'mongodb' => [
            'driver'   => env('MONGODB_DRIVER'),
            'host'     => env('MONGODB_HOST'),
            'port'     => env('MONGODB_PORT'),
            'database' => env('MONGODB_DATABASE'),
            'username' => env('MONGODB_USERNAME'),
            'password' => env('MONGODB_PASSWORD'),
            'options' => [
                'database' => 'admin' // sets the authentication database required by mongo 3
            ]
        ]
    ]
];