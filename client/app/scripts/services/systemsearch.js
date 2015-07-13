'use strict';

/**
 * @ngdoc service
 * @name clientApp.Systems
 * @description
 * # Systems
 * Factory in the clientApp.
 */
angular.module('clientApp')
    .factory('SystemSearch', ['$resource', function($resource) {
        // Service logic

        // Public API here
        return $resource('/api/systems/search/', null, null);
    }]);
