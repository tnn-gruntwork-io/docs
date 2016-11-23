package nav

// These RegEx's help us determine if a given Folder is of a particular type by checking its OutputPath.
const OUTPUT_PATH_IS_PACKAGE_FOLDER_REGEX = `^packages/[\w -]+$`
const OUTPUT_PATH_IS_MODULE_FOLDER_REGEX = `^packages/[\w -]+/modules/[\w -]+$`
