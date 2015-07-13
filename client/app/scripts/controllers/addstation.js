'use strict';

/**
 * @ngdoc function
 * @name clientApp.controller:AddstationCtrl
 * @description
 * # AddstationCtrl
 * Controller of the clientApp
 */
angular.module('clientApp')
    .controller('AddstationCtrl', ['lodash', 'SystemSearch', 'Stations', 'Commodities', 'gameDataService', '$modal', '$rootScope',
        function(lodash, SystemSearch, Stations, Commodities, gameDataService, $modal, $rootScope) {
            var self = this;

            self.difference = lodash.difference;
            self.indexOf = lodash.indexOf;
            self.contains = lodash.contains;

            self.css = {
                labelWidth: '2',
                inputWidth: '9',
                addonWidth: '1'
            };

            self.stationTypes = gameDataService.stationTypes();
            self.stationWealth = gameDataService.stationWealth();
            self.stationPopulation = gameDataService.stationPopulation();
            self.stationEconomies = gameDataService.stationEconomies();
            self.stationFacilities = gameDataService.stationFacilities();
            self.stationAllegiance = gameDataService.allegiances();
            self.stationGovernment = gameDataService.governments();

            self.moduleClasses = gameDataService.moduleClasses();

            self.distUnits = gameDataService.distUnits();

            self.selectedDistUnit = self.distUnits[0];

            // Commodity Entry
            function Snapshot() {
                this.commodity = {
                    name: null,
                    isNew: false,
                };
                this.sellPrice = 0;
                this.buyPrice = 0;
                this.demand = 0;
                this.supply = 0;
                this.galacticAverage = 0;
            }

            self.addSnapshot = function(station) {
                station.snapshots.push(new Snapshot());
            };

            self.snapshotIsValid = function(snapshot) {
                if (lodash.isEmpty(snapshot.commodity) || lodash.isNull(snapshot.commodity.name)) {
                    return false;
                }

                if (typeof snapshot === 'string' && (snapshot.commodity.name.length < 3 ||
                        snapshot.commodity.name.length > 255)) {
                    return false;
                }

                if (lodash.isNull(snapshot.sellPrice) || lodash.isNull(snapshot.buyPrice) ||
                    lodash.isNull(snapshot.demand) || lodash.isNull(snapshot.supply) ||
                    lodash.isNull(snapshot.galacticAverage)) {
                    return false;
                }

                if (snapshot.sellPrice < 0 || snapshot.buyPrice < 0 || snapshot.demand < 0 ||
                    snapshot.supply < 0 || snapshot.galacticAverage < 0) {
                    return false;
                }

                return true;
            };

            // Get a list of valid commodity names.
            self.getCommodityList = function(name) {
                var commodityList = [];
                commodityList = Commodities.list().$promise.then(function(response) {
                    return lodash.first(lodash.filter(response, function(item) {
                        return lodash.contains(item.toLowerCase(), name.toLowerCase());
                    }), 8);
                }, function(errResponse) {
                    console.log(errResponse);
                    return [];
                });

                return commodityList;
            };

            self.commodityExists = function(name) {
                return lodash.contains(self.getCommodityList(), name);
            };

            // Alert messages.
            self.alerts = [];
            self.closeAlert = function(index) {
                self.alerts.splice(index, 1);
            };

            self.findSystem = function(name) {
                var search = new SystemSearch({
                    name: name
                });

                var matches = [];

                matches = search.$save().then(function(response) {
                    if (response.systems === null) {
                        return [];
                    }

                    return response.systems;
                }, function() {
                    return [];
                });

                return matches;
            };

            // Create New Commodities in a Modal Window
            self.createCommodityModal = function(name, galacticAverage, index) {
                var scope = $rootScope.$new();
                scope.params = {
                    name: name,
                    galacticAverage: galacticAverage
                };

                var modalInstance = $modal.open({
                    templateUrl: '/views/addCommodityModal.html',
                    controller: 'AddCommodityModalCtrl as ctrl',
                    scope: scope
                });

                modalInstance.result.then(function(newCommodity) {
                    self.newStation.snapshots[index].commodity.name = newCommodity.name;
                    self.newStation.snapshots[index].commodity.isNew = true;

                    if (newCommodity.galacticAverage) {
                        self.newStation.snapshots[index].galacticAverage = newCommodity.galacticAverage;
                    }
                });
            };

            self.createNewStation = function(newStation) {
                // Calculate distance from jump.
                newStation.distanceFromJump = self.enteredDist * self.selectedDistUnit.factor;

                // Clean Up Object Arrays
                newStation.outfitting = lodash.filter(newStation.outfitting, function(module) {
                    return module && module.name && module.mClass && module.size && module.price;
                });

                newStation.shipyard = lodash.filter(newStation.shipyard, function(ship) {
                    return ship && ship.name && ship.price;
                });

                newStation.snapshots = lodash.filter(newStation.snapshots, function(snapshot) {
                    return self.snapshotIsValid(snapshot);
                });

                // Clean up null values in arrays.
                newStation.outfitting = lodash.compact(newStation.outfitting);
                newStation.shipyard = lodash.compact(newStation.shipyard);
                newStation.snapshots = lodash.compact(newStation.snapshots);

                var station = new Stations(newStation);
                station.$save().then(function(response) {
                        self.alerts.push({
                            type: 'success',
                            message: response.slug
                        });
                    },
                    function(errResponse) {
                        self.alerts.push({
                            type: 'danger',
                            message: errResponse.data
                        });
                    });
            };

            self.reset = function(form) {
                if (form) {
                    form.$setPristine();
                    form.$setUntouched();
                }

                self.newStation = {
                    economies: [null],
                    facilities: [null],
                    outfitting: [{
                        name: null,
                        mClass: null,
                        size: null,
                        price: null
                    }],
                    shipyard: [{
                        name: null,
                        price: null
                    }],
                    snapshots: [new Snapshot()]
                };
            };

            self.reset();
        }
    ]);
