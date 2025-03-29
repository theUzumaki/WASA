<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>

<script>
export default {

	data: function() {
		return {
			picture: null,
			username: null,
			logged: false,
		}
	},
	methods: {
		loadUser(){
			this.logged= true;
			if (sessionStorage.user != null){
				this.picture= JSON.parse(sessionStorage.user).picture
				this.username= JSON.parse(sessionStorage.user).name
			}
		},
		startUserLoading() {
			this.intervalId = setInterval(() => {
        		this.loadUser();
    		}, 1000);
    	},
    	stopUserLoading() {
    		if (this.intervalId) {
        		clearInterval(this.intervalId);
    		}
    	},
	},
	mounted() {
	    this.startUserLoading();
	},
	beforeRouteLeave(){
		this.stopUserLoading();
	},
	beforeUnmount() {
    	this.stopUserLoading();
	},
}
</script>

<template>

	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="#/">Wasa Text</a>
		<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
	</header>

	<div class="container-fluid">
		<div class="row">
			<nav id="sidebarMenu" class="col-md-3 col-lg-2 d-md-block bg-light sidebar collapse">
				<div class="position-sticky pt-3 sidebar-sticky">
					<div v-if=logged class="d-flex align-items-center px-3 mb-3">
						<img :src="picture" alt="User Profile" class="rounded-circle" style="margin-right: 5px;" width="40" height="40">
						{{ username }}
					</div>
					<h6 class="sidebar-heading d-flex justify-content-between align-items-center px-3 mb-1 text-muted text-uppercase">
						<span>General</span>
					</h6>
					<ul class="nav flex-column">
						<li class="nav-item">
							<RouterLink to="/home" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#home"/></svg>
								home
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/settings" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#layout"/></svg>
								settings
							</RouterLink>
						</li>
						<li class="nav-item">
							<RouterLink to="/" class="nav-link">
								<svg class="feather"><use href="/feather-sprite-v4.29.0.svg#key"/></svg>
								logout/login
							</RouterLink>
						</li>
					</ul>
				</div>
			</nav>

		</div>
	</div>
	<main class="col-md-9 ms-sm-auto col-lg-10 px-md-4">
		<RouterView />
	</main>
</template>

<style>

.main {
	margin-top: 95px;
	margin: 0 auto;
	padding: 1em;
}

.homescreen {
	position: absolute;
	top: 19%;
	left: 18%;
	bottom: 0px;
	right: 0px;
}

.btn-to-the-right {
	position: relative;
	top: 0px;
}

</style>
