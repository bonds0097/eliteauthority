'use strict';

/**
 * @ngdoc overview
 * @name clientApp
 * @description
 * # clientApp
 *
 * Main module of the application.
 */
angular
    .module('clientApp', [
        'ngRoute',
        'ngResource',
        'ui.bootstrap',
        'ngLodash',
        'ui.grid'
    ])
    .config(['$routeProvider', '$locationProvider', '$resourceProvider',
        function($routeProvider, $locationProvider, $resourceProvider) {
            $routeProvider
                .when('/', {
                    templateUrl: 'views/main.html',
                    controller: 'MainCtrl as ctrl'
                })
                .when('/stations/add', {
                    templateUrl: 'views/addstation.html',
                    controller: 'AddstationCtrl as ctrl'
                })
                .when('/stations', {
                    templateUrl: 'views/stations.html',
                    controller: 'StationsCtrl as ctrl'
                })
                .when('/stations/:slug', {
                    templateUrl: 'views/viewStation.html',
                    controller: 'ViewStationCtrl as ctrl'
                })
                .otherwise({
                    redirectTo: '/'
                });

            // Don't strip trailing slashes from calculated URLs
            $resourceProvider.defaults.stripTrailingSlashes = false;

            // use the HTML5 History API
            $locationProvider.html5Mode(true);
        }
    ]);
