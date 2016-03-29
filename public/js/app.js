(function() {

	var app = angular.module('app', []);

	app.directive('loginForm', function() {
		return {
			restrict: 'E',
			templateUrl: 'html/login.view.html',
			controller: function() {

			},
			controllerAs: 'panels'
		}
	});

})();