'use strict';

/**
 * @ngdoc service
 * @name clientApp.gameDataService
 * @description
 * # gameDataService
 * Factory in the clientApp.
 */
angular.module('clientApp')
    .factory('gameDataService', function() {
        // Service logic
        var commodityCategories = ['Chemicals', 'Consumer Items', 'Foods', 'Industrial Materials',
            'Legal Drugs', 'Machinery', 'Medicines', 'Metals', 'Minerals', 'Salvage', 'Slavery',
            'Technology', 'Textiles', 'Waste', 'Weapons'
        ];

        var stationTypes = ['Coriolis', 'Orbis', 'Ocellus', 'Asteroid', 'O\'Neill Cylinder'];
        var stationWealth = ['Poor', 'Normal', 'Rich'];
        var stationPopulation = ['Small', 'Normal', 'Large'];
        var stationEconomies = ['Agriculture', 'Extraction', 'Industrial', 'High Tech', 'Military',
            'Refinery', 'Service', 'Terraforming', 'Tourism'
        ];
        var stationFacilities = ['Commodities', 'Refuel', 'Repair', 'Re-arm', 'Outfitting', 'Shipyard'];

        var allegiances = ['Empire', 'Independent', 'Federation', 'Alliance'];
        var governments = ['Anarchy', 'Colony', 'Communism', 'Confederacy', 'Cooperative',
            'Corporate', 'Democracy', 'Dictatorship', 'Feudal', 'Imperial', 'Patronage',
            'Prison Colony', 'Theocracy'
        ];

        var moduleClasses = ['E', 'D', 'C', 'B', 'A'];

        var distUnits = [{
            label: 'Ls',
            factor: 1
        }, {
            label: 'Lm',
            factor: 60
        }, {
            label: 'Lh',
            factor: 3600
        }, {
            label: 'Ld',
            factor: 86400
        }, {
            label: 'LY',
            factor: 31556926
        }];

        // Public API here
        return {
            commodityCategories: function() {
                return commodityCategories;
            },
            stationTypes: function() {
                return stationTypes;
            },
            stationWealth: function() {
                return stationWealth;
            },
            stationPopulation: function() {
                return stationPopulation;
            },
            stationEconomies: function() {
                return stationEconomies;
            },
            stationFacilities: function() {
                return stationFacilities;
            },
            allegiances: function() {
                return allegiances;
            },
            governments: function() {
                return governments;
            },
            moduleClasses: function() {
                return moduleClasses;
            },
            distUnits: function() {
                return distUnits;
            }
        };
    });
