<?php

namespace App\Models;

use Jenssegers\Mongodb\Eloquent\Model as Eloquent;

class Heat extends Eloquent
{
	protected $connection = 'mongodb';
    protected $collection = 'usersheats';
    protected $primaryKey = '_id';
    protected $fillable = [
        '_id',
        'user',
        'x',
        'y'
    ];
}