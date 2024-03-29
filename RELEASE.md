# Releasing ManageF

This document gives the steps required to perform a release of ManageF. This document is only useful to ManageF release managers.

## Steps

The steps to perform a release are very simple and easy to execute thanks to the way [Travis](.travis.yml) is setup.
Before running these steps, it is assumed the [master](https://github.com/managef/api/tree/master) branch is all up to date and includes all code that is to be released.

1) Update the [makefile](Makefile) so the VERSION variable is set to the new version number of the upcoming release. For example,

```
VERSION ?= 1.0.0.Final
```

2) Commit that change to the `master` branch and push that change to the GitHub repository:

```
$ git add Makefile
$ git commit -m "New release"
$ git push origin master
```

NOTE: At this point, Travis will run the test suite, build the binary, and push the "latest" docker images to [DockerHub](https://hub.docker.com/r/aljesusg/managef_api/).

3) Once Travis is "green" (which confirms all tests have passed and the latest images have been pushed to DockerHub), you can now tag the release. For example,

```
$ git tag 1.0.0.Final
$ git push --tags
```

NOTE: Once the tag has been pushed, Travis will immediately run another build but this time will tag the docker images with the same version string as the git tag name.
In the example above, the docker images pushed to DockerHub will not be tagged with "latest" but will instead be tagged with "1.0.0.Final."

## Prepare For The Next Release

Now that the release is done, you will want to bump up the version with a new version so any future master ("latest") builds will be given the new future version. Preferrably you should use a version with "SNAPSHOT" in it to indicate it is not a release but instead a work in progress (e.g. `1.0.1.Final-SNAPSHOT`). So simply repeat step (1) with the snapshot version string.

## Finished Product

Once the above steps have been performed, the release is complete.
There is now a git tag that marks the code that produced the release and there are DockerHub images for that release.

* [Git tags](https://github.com/managef/api/tags)
* DockerHub images:
   * [ManageF-API](https://hub.docker.com/r/aljesusg/managef_api/tags/)

## How It Works

Most of the magic is set up in the [travis](.travis.yml) file.

If the Travis build is a tag build, then the tag name is passed
to all Makefiles as the value of the DOCKER_VERSION environment variable.
The Makefiles use that to tag the docker images when building and pushing the images to DockerHub.

If the Travis build is a branch build (that is, not triggered by a tag), the name of the branch is set to the DOCKER_VERSION variable
with the exception of the `master` branch. If the branch name is `master`, the DOCKER_VERSION is set to "latest".
Again, the value of the DOCKER_VERSION variable is then used by the Makefiles to tag the docker images.

Note that only tags in the format `\#.#.#[.Label]` will trigger tag builds and, today, only the `master` branch will trigger a branch build.
If a release build is to be built off of a non-master branch, simply add the branch name to the "branches" section of the .travis.yml file.

