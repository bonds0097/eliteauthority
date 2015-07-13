'use strict';

/**
 * @ngdoc function
 * @name clientApp.controller:StationsCtrl
 * @description
 * # StationsCtrl
 * Controller of the clientApp
 */
angular.module('clientApp')
    .controller('StationsCtrl', ['$location', 'lodash', 'Stations', function($location, lodash, Stations) {
        var self = this;

        self.recentStations = Stations.query();

        self.loadStation = function(slug) {
            $location.path('/stations/' + slug);
        };

        self.generateStationDescription = function(station) {
            var description = '';

            if (station.population !== 'Normal') {
                description += station.population + ' ';
            }

            if (station.wealth !== 'Normal') {
                description += station.wealth + ' ';
            }

            station.economies.sort().reverse();
            for (var econRemaining = station.economies.length; econRemaining > 0; econRemaining--) {
                description += station.economies[econRemaining - 1];
                if (econRemaining > 2) {
                    description += ', ';
                } else if (econRemaining > 1) {
                    description += ' and ';
                }
            }

            return description.trim();
        };
    }]);
