<script lang="ts">
    import { page } from '$app/stores'
    import Icon from '$lib/Icon.svelte'
    import { DropdownMenu } from '$lib/wildcard'
    import { mdiDotsHorizontal } from '@mdi/js'

    import SidebarToggleButton from './SidebarToggleButton.svelte'
    import { sidebarOpen } from './stores'
    import { navFromPath } from './utils'
    import { getButtonClassName } from '$lib/wildcard/Button'
    import { overflow } from '$lib/dom'

    $: breadcrumbs = navFromPath($page.params.path, $page.params.repo)
</script>

<div class="header">
    <div class="toggle-wrapper" class:hidden={$sidebarOpen}>
        <SidebarToggleButton />
    </div>
    <h2>
        {#each breadcrumbs as [name, path], index}
            <!--
                The elements are arranged like this because we want to
                ensure that the leading / before a segement always stay with
                the segment. I.e.

                    path / to / file

                is wrapped as

                    path
                    / to
                    / file

                However, without the following space the path wouldn't break/wrap
                at all.
            -->
            {' '}
            <span class:last={index === breadcrumbs.length - 1}>
                {#if index > 0}
                    /
                {/if}
                {#if path}
                    <a href={path}>{name}</a>
                {:else}
                    <slot name="icon" />
                    {name}
                {/if}
            </span>
        {/each}
    </h2>
    <div class="actions" use:overflow={{ class: 'compact', measureClass: 'measure' }}>
        <slot name="actions" />
        {#if $$slots.actionmenu}
            <div class="divider" />
            <div class="more">
                <DropdownMenu
                    triggerButtonClass={getButtonClassName({ variant: 'icon' })}
                    aria-label="Show more actions"
                >
                    <svelte:fragment slot="trigger">
                        <Icon svgPath={mdiDotsHorizontal} inline />
                    </svelte:fragment>
                    <slot name="actionmenu" />
                </DropdownMenu>
            </div>
        {/if}
    </div>
</div>

<style lang="scss">
    .header {
        display: flex;
        align-items: baseline;
        padding: 0.25rem 0.5rem;
        border-bottom: 1px solid var(--border-color);
        background-color: var(--color-bg-1);
        z-index: 1;
        gap: 0.5rem;
    }

    .toggle-wrapper {
        margin-right: 0.5rem;

        // We still want the height of the button to be considered
        // when rendering the header, so that toggling the sidebar
        // won't change the height of the header.
        &.hidden {
            visibility: hidden;
            margin-right: 0;
            width: 0;
        }
    }

    h2 {
        font-weight: 400;
        font-size: var(--code-font-size);
        font-family: var(--code-font-family);
        margin: 0;

        a {
            color: var(--body-color);
        }

        span {
            white-space: nowrap;
        }

        .last {
            font-weight: bold;
        }
    }

    .actions {
        --color: var(--icon-color);

        // Ensures that the actions are centered vertically when the header expands
        // due to path name wrapping.
        align-self: center;

        // In order to hide action labels when necessary (i.e. when there wouldn't be
        // enough space to display the path). We use the `overflow` action to measure
        // the space available for the actions. For this to work we need to setup th
        // CSS rules accordingly.

        // Ensures that the element takes up as much space as available.
        flex: 1;
        // Allows the element to shrink past it's content size. This allows us, together
        // with the .measure CSS class, to determine whether the actions would use more
        // space if the labels are visible.
        overflow: hidden;
        // Due to flex: 1 the element will take up all available space, but we want
        // the actions to appear as far to the right as possible.
        justify-content: right;

        // Here is how this works togther: The header starts out with enough space.
        // The actions take up all remaining space, due to `flex: 1`.
        //
        //   ┌──────────────────┐┌──────────────────────────────┐
        //   │ path / to / file ││          [a] Label [b] Label │
        //   └──────────────────┘└──────────────────────────────┘
        //
        // As the header shrinks, the actions will use less space and eventually the content
        // will be hidden due to `overflow: hidden`.
        //
        //   ┌──────────────────────────────┐┌──────────────────┐
        //   │ path / to / file             ││] Label [b] Label │
        //   └──────────────────────────────┘└──────────────────┘
        //
        // At this point the overflow action will "trigger" and compare the current size of the
        // actions element with the size after applying the `.measure` class. The measure class
        // Removes the `overflow: hidden` rule, to prevent the actions element to shrink past
        // its content size. Therefore the overflow action measures the following:
        //
        //                                    (without .measure)
        //   ┌──────────────────────────────┐┌──────────────────┐
        //   │ path / to / file             ││] Label [b] Label │
        //   └──────────────────────────────┘└──────────────────┘
        //                                 .measure
        //   ┌───────────────────────────┐┌─────────────────────┐
        //   │ path / to / file          ││ [a] Label [b] Label │
        //   └───────────────────────────┘└─────────────────────┘
        //
        // It determines that the actions element would use more space if fully displayed and
        // therefore adds the `.compact` class to the actions element, which hides the labels.
        //
        //                                    .compact
        //   ┌──────────────────────────────┐┌──────────────────┐
        //   │ path / to / file             ││          [a] [b] │
        //   └──────────────────────────────┘└──────────────────┘
        //

        // To make the actions menu button appear visually "centered"
        margin-right: 0.5rem;

        display: flex;
        gap: 1rem;
        align-items: center;

        // With overflow:visible the actions won't shrink past their content size,
        // and this allows us to measure the space needed to show actions fully.
        &:global(.measure) {
            overflow: visible;
        }

        // When the actions are "compact" we hide the labels.
        &:global(.compact) {
            // This is necessary to prevent shrinking the actions even past its
            // "compact" size.
            overflow: visible;

            :global([data-action-label]) {
                display: none;
            }
        }

        .divider {
            border-left: 1px solid var(--border-color);
            align-self: stretch;
        }

        .more {
            align-self: center;
        }
    }
</style>
