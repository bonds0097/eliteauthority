'use strict';

angular.module('clientApp')
    .factory('Commodities', ['$resource', function($resource) {
        return $resource('/api/commodities/', null, {
            list: {
                method: 'GET',
                url: '/api/commodities/list/',
                isArray: true,
                cache: false
            }
        });
    }]);
