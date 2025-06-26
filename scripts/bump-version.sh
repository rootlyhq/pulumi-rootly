#!/bin/bash

# Version bumping script for pulumi-rootly
# Usage: ./bump-version.sh [patch|minor|major|show] [patch|minor|major]

set -euo pipefail

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Function to print colored output
print_info() {
    echo -e "${BLUE}ℹ️  $1${NC}"
}

print_success() {
    echo -e "${GREEN}✅ $1${NC}"
}

print_warning() {
    echo -e "${YELLOW}⚠️  $1${NC}"
}

print_error() {
    echo -e "${RED}❌ $1${NC}"
}

# Function to get current version from git tags
get_current_version() {
    local version
    version=$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0")
    echo "$version"
}

# Function to parse version into components
parse_version() {
    local version="$1"
    # Remove 'v' prefix if present
    version="${version#v}"
    
    # Split version into major.minor.patch
    IFS='.' read -ra VERSION_PARTS <<< "$version"
    
    if [[ ${#VERSION_PARTS[@]} -ne 3 ]]; then
        print_error "Invalid version format: $version. Expected format: x.y.z"
        exit 1
    fi
    
    MAJOR="${VERSION_PARTS[0]}"
    MINOR="${VERSION_PARTS[1]}"
    PATCH="${VERSION_PARTS[2]}"
    
    # Validate that all parts are numbers
    if ! [[ "$MAJOR" =~ ^[0-9]+$ ]] || ! [[ "$MINOR" =~ ^[0-9]+$ ]] || ! [[ "$PATCH" =~ ^[0-9]+$ ]]; then
        print_error "Version parts must be numbers: $version"
        exit 1
    fi
}

# Function to calculate next version
calculate_next_version() {
    local bump_type="$1"
    local current_version="$2"
    
    parse_version "$current_version"
    
    case "$bump_type" in
        "patch")
            echo "$MAJOR.$MINOR.$((PATCH + 1))"
            ;;
        "minor")
            echo "$MAJOR.$((MINOR + 1)).0"
            ;;
        "major")
            echo "$((MAJOR + 1)).0.0"
            ;;
        *)
            print_error "Invalid bump type: $bump_type. Use: patch, minor, or major"
            exit 1
            ;;
    esac
}

# Function to check git status
check_git_status() {
    if ! git diff-index --quiet HEAD --; then
        print_error "Working directory is not clean. Please commit or stash your changes."
        git status --short
        exit 1
    fi
    
    # Check if we're on the main/master branch
    local current_branch
    current_branch=$(git rev-parse --abbrev-ref HEAD)
    if [[ "$current_branch" != "main" ]] && [[ "$current_branch" != "master" ]]; then
        print_warning "Not on main/master branch (currently on: $current_branch)"
        read -p "Continue anyway? (y/N): " -r
        if [[ ! $REPLY =~ ^[Yy]$ ]]; then
            print_info "Aborting version bump."
            exit 0
        fi
    fi
}

# Function to create and push git tag
create_and_push_tag() {
    local version="$1"
    local tag="v$version"
    
    print_info "Creating git tag: $tag"
    git tag -a "$tag" -m "Release $tag"
    
    print_info "Pushing tag to origin..."
    git push origin "$tag"
    
    print_success "Tag $tag created and pushed successfully!"
}

# Function to show version information
show_version() {
    local current_version
    current_version=$(get_current_version)
    
    echo "Current version: $current_version"
    
    if [[ $# -gt 1 ]]; then
        local bump_type="$2"
        local next_version
        next_version=$(calculate_next_version "$bump_type" "$current_version")
        echo "$next_version"
    else
        echo "Next patch: $(calculate_next_version "patch" "$current_version")"
        echo "Next minor: $(calculate_next_version "minor" "$current_version")"
        echo "Next major: $(calculate_next_version "major" "$current_version")"
    fi
}

# Function to bump version
bump_version() {
    local bump_type="$1"
    local current_version
    local next_version
    
    current_version=$(get_current_version)
    next_version=$(calculate_next_version "$bump_type" "$current_version")
    
    print_info "Current version: $current_version"
    print_info "Next $bump_type version: v$next_version"
    
    # Confirm with user
    read -p "Proceed with $bump_type version bump? (y/N): " -r
    if [[ ! $REPLY =~ ^[Yy]$ ]]; then
        print_info "Aborting version bump."
        exit 0
    fi
    
    # Check git status
    check_git_status
    
    # Create and push tag
    create_and_push_tag "$next_version"
}

# Function to show help
show_help() {
    echo "Version Bumping Script for pulumi-rootly"
    echo ""
    echo "Usage:"
    echo "  $0 patch                    # Bump patch version (1.2.3 → 1.2.4)"
    echo "  $0 minor                    # Bump minor version (1.2.3 → 1.3.0)"
    echo "  $0 major                    # Bump major version (1.2.3 → 2.0.0)"
    echo "  $0 show [patch|minor|major] # Show current/next version(s)"
    echo "  $0 help                     # Show this help message"
    echo ""
    echo "Examples:"
    echo "  $0 patch                    # Bump patch and push tag"
    echo "  $0 show                     # Show all version info"
    echo "  $0 show minor               # Show next minor version only"
    echo ""
    echo "Notes:"
    echo "  • This script uses git tags for version management"
    echo "  • Working directory must be clean before bumping"
    echo "  • Tags are automatically pushed to origin"
    echo "  • CI will trigger releases when tags are pushed"
    echo ""
}

# Main script logic
main() {
    if [[ $# -eq 0 ]]; then
        show_help
        exit 1
    fi
    
    case "$1" in
        "patch"|"minor"|"major")
            bump_version "$1"
            ;;
        "show")
            show_version "$@"
            ;;
        "help"|"-h"|"--help")
            show_help
            ;;
        *)
            print_error "Invalid command: $1"
            show_help
            exit 1
            ;;
    esac
}

# Check if git is available
if ! command -v git &> /dev/null; then
    print_error "Git is not installed or not in PATH"
    exit 1
fi

# Check if we're in a git repository
if ! git rev-parse --git-dir > /dev/null 2>&1; then
    print_error "Not in a git repository"
    exit 1
fi

# Run main function
main "$@"