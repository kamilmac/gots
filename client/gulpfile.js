var gulp = require('gulp'),
    postcss = require('gulp-postcss'),
    cssnext = require('cssnext'),
    atImport = require('postcss-import'),
    cssmin = require('gulp-cssmin'),
    rename = require('gulp-rename'),
    livereload = require('gulp-livereload'),
    runSeq = require('run-sequence'),
    processors = [
        atImport(),
        cssnext()
    ];

var vendor = [
    "./node_modules/vue/dist/vue.min.js",
    "./node_modules/whatwg-fetch/fetch.js",
    "./node_modules/systemjs/dist/system.js"
];
 
gulp.task('css:bundle', function () {   
    return gulp.src('./static/css/main.css')
        .pipe(postcss(processors))
        .pipe(rename(function (path) {
            path.basename = "bundle";
        }))
        .pipe(cssmin())
        .pipe(gulp.dest('./static/css'));
});

gulp.task('js:bundle', function () {
    var merge = require('merge2'),
        ts = require('gulp-typescript'),
        uglify = require('gulp-uglify'),
        tscConfig = ts.createProject('./static/tsconfig.json'),
        concat = require('gulp-concat');
        
	var vendorBundle = gulp.src(vendor)
		.pipe(concat("vendor.js"));
    var tsBundle = gulp.src('./static/js/**/*.ts')
		.pipe(ts(tscConfig))
        .pipe(uglify());
        
        
        
    return merge(vendorBundle, tsBundle)
        .pipe(concat("bundle.js"))
		.pipe(gulp.dest('./static/js'));
});

gulp.task('reload', function() {
    gulp.src('')
        .pipe(livereload());
});

gulp.task('css:watch', function() {
    gulp.watch([
        './static/css/**/*.css',
        '!./static/css/bundle.css'
    ], function() { runSeq('css:bundle', 'reload') });
});

gulp.task('js:watch', function() {
    gulp.watch([
        './static/js/**/*.ts'
    ], ['reload']);
});

gulp.task('watch', ['css:watch', 'js:watch'], function() {
    livereload.listen();
});