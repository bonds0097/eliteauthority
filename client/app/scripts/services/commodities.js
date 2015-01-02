'use strict';

angular.module('clientApp')
    .factory('Commodities', ['$resource', function($resource) {
        return $resource('/api/commodities/:slug', null);
    }]);
