'use strict';

/**
 * @ngdoc service
 * @name clientApp.stations
 * @description
 * # stations
 * Factory in the clientApp.
 */
angular.module('clientApp')
    .factory('Stations', ['$resource', function($resource) {

        // Public API here
        return $resource('/api/stations/:slug', null, {
            update: {
                method: 'PUT'
            }
        });
    }]);
