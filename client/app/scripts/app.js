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
        'ui.bootstrap'
    ])
    .config(['$routeProvider', '$locationProvider', '$resourceProvider',
        function($routeProvider, $locationProvider, $resourceProvider) {
            $routeProvider
                .when('/', {
                    templateUrl: 'views/main.html',
                    controller: 'MainCtrl as ctrl'
                })
                .when('/second', {
                    templateUrl: 'views/second.html',
                    controller: 'SecondCtrl as ctrl'
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
