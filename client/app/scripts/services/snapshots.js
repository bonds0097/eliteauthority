'use strict';

angular.module('clientApp')
    .factory('Snapshots', ['$resource', function($resource) {
        return $resource('/api/snapshots/', null);
    }]);
