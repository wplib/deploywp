################################################################################
# Workflow summary


################################################################################
# GoLang
rm deploy/*	# plus any other . files.
cp src/composer.json deploy
# Modify composer.json based on git branch.

src/{{ .extra.wordpress-webroot-path }}/{{ .extra.wordpress-core-path }}/ -> deploy/
src/{{ .extra.wordpress-webroot-path }}/{{ .extra.wordpress-content-path }}/ -> deploy/wp-content/


################################################################################
# Shell
composer install
rm src/composer.*


################################################################################
# Shell
git add .

git commit -m 'blah' .

git push

